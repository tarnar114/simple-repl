package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type OperationType int64

const (
	META_CMD_UNRECOGNIZED_CMD OperationType = iota
	META_CMD_SUCCESS
	PREPARE_UNRECOGNIZED_RESULT
	PREPARE_SUCCESS
)

func (s OperationType) String() string {
	switch s {
	case META_CMD_UNRECOGNIZED_CMD:
		return "unknown meta_cmd statement"
	case META_CMD_SUCCESS:
		return "meta_cmd success"
	case PREPARE_UNRECOGNIZED_RESULT:
		return "prepare_cmd unrecognized"
	case PREPARE_SUCCESS:
		return "prepare_cmd success"
	}
	return "unknown"
}

type StatementType int

const (
	INSERT_STATEMENT StatementType = iota
	SELECT_STATEMENT
)

type Statement struct {
	Type StatementType
}

func (s StatementType) String() string {
	switch s {
	case INSERT_STATEMENT:
		return "insert statement"
	case SELECT_STATEMENT:
		return "select statement"
	}
	return "unknown"
}

//first time prompt
func prompt() {
	fmt.Print("go-sqlite> ")
}

//cleans input
func InputClean(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

//returns help prompt
func displayHelp() {
	fmt.Println("Welcome to goSQLite! These are the available commands: ")
	fmt.Println("\t .help    - Show available commands")
	fmt.Println("\t .clear   - Clear the terminal screen")
	fmt.Println("\t .exit    - Closes your connection to database")
}

//cmd to clear
func clrScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func exitCLI() {
	os.Exit(0)
}

//func to handle meta cmds
func handleMetaCmd(text string) OperationType {
	commands := map[string]interface{}{
		".help":  displayHelp,
		".clear": clrScreen,
		".exit":  exitCLI,
	}
	if command, exists := commands[text]; exists {
		command.(func())()
		return META_CMD_SUCCESS
	} else {
		return META_CMD_UNRECOGNIZED_CMD
	}
}

func handlePrepStatement(text string, statement *Statement) OperationType {
	if strings.EqualFold(text, "insert") {
		statement.Type = INSERT_STATEMENT
		return PREPARE_SUCCESS
	} else if strings.EqualFold(text, "select") {
		statement.Type = SELECT_STATEMENT
		return PREPARE_SUCCESS
	}
	return PREPARE_UNRECOGNIZED_RESULT
}
func execStatement(statement Statement) {
	switch statement.Type {
	case (SELECT_STATEMENT):
		fmt.Println(SELECT_STATEMENT)
	case (INSERT_STATEMENT):
		fmt.Println(INSERT_STATEMENT)
	}
}
func main() {
	reader := bufio.NewScanner(os.Stdin)

	prompt()
	for reader.Scan() {
		text := InputClean(reader.Text())
		if text[0:1] == "." {
			switch handleMetaCmd(text) {
			case (META_CMD_SUCCESS):
				prompt()
				continue
			case (META_CMD_UNRECOGNIZED_CMD):
				fmt.Println(META_CMD_UNRECOGNIZED_CMD)
				prompt()
				continue
			}

		}
		var statement Statement

		switch handlePrepStatement(text, &statement) {
		case (PREPARE_SUCCESS):
			execStatement(statement)
			prompt()
		case (PREPARE_UNRECOGNIZED_RESULT):
			println("unrecognized statement: %s", text)
			prompt()
			continue
		}

	}
	fmt.Println()
}
