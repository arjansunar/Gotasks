# Gotasks

Todo tracker cli

> This project is done as a part of [task-tracker](https://roadmap.sh/projects/task-tracker) from [roadmap.sh](https://roadmap.sh/)

## How to Run

```bash
git clone https://github.com/arjansunar/gotasks
cd gotasks
go build -o gotasks
```

### Installation

```bash
go install github.com/arjansunar/gotasks
```

## Usage

```bash
./gotasks help

```

```txt
Usage: gotasks <command> [options]

Commands:
add <task_description> Add a new task
update <task_id> <new_description> Update an existing task
delete <task_id> Delete a task
mark-in-progress <task_id> Mark a task as in progress
mark-done <task_id> Mark a task as done
list List all tasks
list <status> List tasks by status (done, todo, in-progress)
help Show this message and exit

Examples:
gotasks add "Buy groceries"
gotasks update 1 "Buy groceries and cook dinner"
gotasks delete 1
gotasks mark-in-progress 1
gotasks mark-done 1
gotasks list
gotasks list done
gotasks list todo
gotasks list in-progress
```
