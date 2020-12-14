// Copyright 2017 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package casbin

import (
	"encoding/json"
	"errors"
	"runtime"

	"github.com/casbin/casbin/model"
	"github.com/casbin/casbin/persist"
	"github.com/joshua-chen/go-commons/middleware/perm"
	"github.com/kataras/golog"
	"github.com/lib/pq"
	"github.com/xormplus/xorm"

)

// Adapter represents the Xorm adapter for policy storage.
type Adapter struct {
	driverName     string
	dataSourceName string
	dbSpecified    bool
	engine         *xorm.Engine
}

// finalizer is the destructor for Adapter.
func finalizer(a *Adapter) {
	a.engine.Close()
}

// NewAdapter is the constructor for Adapter.
// dbSpecified is an optional bool parameter. The default value is false.
// It's up to whether you have specified an existing DB in dataSourceName.
// If dbSpecified == true, you need to make sure the DB in dataSourceName exists.
// If dbSpecified == false, the adapter will automatically create a DB named "casbin".
func NewAdapter(driverName string, dataSourceName string, dbSpecified ...bool) *Adapter {
	a := &Adapter{}
	a.driverName = driverName
	a.dataSourceName = dataSourceName

	if len(dbSpecified) == 0 {
		a.dbSpecified = false
	} else if len(dbSpecified) == 1 {
		a.dbSpecified = dbSpecified[0]
	} else {
		panic(errors.New("invalid parameter: dbSpecified"))
	}

	// Open the DB, create it if not existed.
	a.open()

	// Call the destructor when the object is released.
	runtime.SetFinalizer(a, finalizer)

	return a
}

func (a *Adapter) createDatabase() error {
	var err error
	var engine *xorm.Engine
	if a.driverName == "postgres" {
		engine, err = xorm.NewEngine(a.driverName, a.dataSourceName+" dbname=postgres")
	} else {
		engine, err = xorm.NewEngine(a.driverName, a.dataSourceName)
	}
	if err != nil {
		return err
	}
	defer engine.Close()

	if a.driverName == "postgres" {
		if _, err = engine.Exec("CREATE DATABASE casbin"); err != nil {
			// 42P04 is	duplicate_database
			if pqerr, ok := err.(*pq.Error); ok && pqerr.Code == "42P04" {
				return nil
			}
		}
	} else {
		_, err = engine.Exec("CREATE DATABASE IF NOT EXISTS casbin")
	}
	return err
}

func (a *Adapter) open() {
	var err error
	var engine *xorm.Engine

	if a.dbSpecified {
		engine, err = xorm.NewEngine(a.driverName, a.dataSourceName)
		if err != nil {
			panic(err)
		}
	} else {
		if err = a.createDatabase(); err != nil {
			panic(err)
		}

		if a.driverName == "postgres" {
			engine, err = xorm.NewEngine(a.driverName, a.dataSourceName+" dbname=casbin")
		} else {
			engine, err = xorm.NewEngine(a.driverName, a.dataSourceName+"casbin")
		}
		if err != nil {
			panic(err)
		}
	}

	a.engine = engine

	a.createTable()
}

func (a *Adapter) close() {
	a.engine.Close()
	a.engine = nil
}

func (a *Adapter) createTable() {
	err := a.engine.Sync2(new(perm.CasbinRule))
	if err != nil {
		panic(err)
	}
}

func (a *Adapter) dropTable() {
	err := a.engine.DropTables(new(perm.CasbinRule))
	if err != nil {
		panic(err)
	}
}

func loadPolicyLine(line perm.CasbinRule, model model.Model) {
	lineText := line.PType
	if line.Sub != "" {
		lineText += ", " + line.Sub
	}
	if line.Obj != "" {
		lineText += ", " + line.Obj
	}
	if line.Act != "" {
		lineText += ", " + line.Act
	}
	if line.Suf != "" {
		lineText += ", " + line.Suf
	}
	/*
		if line.Name != "" {
			lineText += ", " + line.Name
		}
		if line.Des != "" {
			lineText += ", " + line.Des
		}

	*/

	persist.LoadPolicyLine(lineText, model)
}

// LoadPolicy loads policy from database.
func (a *Adapter) LoadPolicy(model model.Model) error {
	var lines []perm.CasbinRule
	err := a.engine.Find(&lines)
	if err != nil {
		return err
	}

	for _, line := range lines {
		loadPolicyLine(line, model)
	}

	jsonData, _ := json.Marshal(lines)
	golog.Debugf("CasbinRule model %s: ", string(jsonData))

	return nil
}

func savePolicyLine(ptype string, rule []string) perm.CasbinRule {
	line := perm.CasbinRule{}

	line.PType = ptype
	if len(rule) > 0 {
		line.Sub = rule[0]
	}
	if len(rule) > 1 {
		line.Obj = rule[1]
	}
	if len(rule) > 2 {
		line.Act = rule[2]
	}
	if len(rule) > 3 {
		line.Suf = rule[3]
	}
	if len(rule) > 4 {
		line.Name = rule[4]
	}
	if len(rule) > 5 {
		line.Des = rule[5]
	}

	return line
}

// SavePolicy saves policy to database.
func (a *Adapter) SavePolicy(model model.Model) error {
	a.dropTable()
	a.createTable()

	var lines []perm.CasbinRule

	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			lines = append(lines, line)
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			lines = append(lines, line)
		}
	}

	_, err := a.engine.Insert(&lines)
	return err
}

// AddPolicy adds a policy rule to the storage.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	_, err := a.engine.Insert(line)
	return err
}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	_, err := a.engine.Delete(line)
	return err
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	line := perm.CasbinRule{}

	line.PType = ptype
	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		line.Sub = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		line.Obj = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		line.Act = fieldValues[2-fieldIndex]
	}
	if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) {
		line.Suf = fieldValues[3-fieldIndex]
	}
	if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) {
		line.Name = fieldValues[4-fieldIndex]
	}
	if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) {
		line.Des = fieldValues[5-fieldIndex]
	}

	_, err := a.engine.Delete(line)
	return err
}
