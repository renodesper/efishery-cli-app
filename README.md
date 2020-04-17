# eFishery CLI App

This project serves as my solution for eFishery test project, which is an offline-first CLI app to manage tasks. Please make sure to use a Golang version which already supports Go modules.

## How to build

To build this project, I already provide a Makefile to ease the building. After the build, there will be a new file called `efishery-cli-app`.

```sh
make
```

## Available Commands

To check the available commands, we can call the help of the app with `-h` flag.

```sh
❯ ./efishery-cli-app -h
A brief description of your application

Usage:
  efishery-cli-app [command]

Available Commands:
  add         Add a new task
  delete      Delete a task
  done        Set specific task as done
  edit        Edit specific task
  help        Help about any command
  list        Show all tasks (offline-first, call remote database when connection is available)
  sync        Sync local data into remote database

Flags:
      --config string    (default "efishery-cli-app.toml")
  -h, --help            help for efishery-cli-app

Use "efishery-cli-app [command] --help" for more information about a command.
```

### Add Command

```sh
❯ ./efishery-cli-app add -h
Add a new task

Usage:
  efishery-cli-app add [flags]

Flags:
  -c, --content string   Task content, wrap text with quotation mark
  -h, --help             help for add
  -t, --tags string      Tags for the task, separated by comma
```

### Delete Command

```sh
❯ ./efishery-cli-app delete -h
Delete a task

Usage:
  efishery-cli-app delete [flags]

Flags:
  -h, --help        help for delete
  -i, --id string   docID of the task
```

### Done Command

```sh
❯ ./efishery-cli-app done -h
Set specific task as done

Usage:
  efishery-cli-app done [flags]

Flags:
  -h, --help        help for done
  -i, --id string   docID of the task
```

### Edit Command

```sh
❯ ./efishery-cli-app edit -h
Edit specific task

Usage:
  efishery-cli-app edit [flags]

Flags:
  -c, --content string   Task content, wrap text with quotation mark
  -h, --help             help for edit
  -i, --id string        docID of the task
  -t, --tags string      Tags for the task, separated by comma
```

### List Command

```sh
❯ ./efishery-cli-app list -h
Show all tasks (offline-first, call remote database when connection is available)

Usage:
  efishery-cli-app list [flags]

Flags:
  -h, --help   help for list
```

### Sync Command

```sh
❯ ./efishery-cli-app sync -h
Sync local data into remote database

Usage:
  efishery-cli-app sync [flags]

Flags:
  -h, --help   help for sync
```

## Example

On this example, I assume that all commands will be executed one by one to make sure the result can be used by the next command.

Let's add some tasks first.

```sh
❯ ./efishery-cli-app add -c "Content 1" -t tag1
❯ ./efishery-cli-app add -c "Content 2" -t tag2
❯ ./efishery-cli-app add -c "Content 3" -t tag3
```

Check whether the data is exists or not.

```sh
❯ ./efishery-cli-app list

docID                                   Status  Tags    Content
d63c76b0-80b8-11ea-983a-68f72850e714    Todo    tag1    Content 1
dd6badac-80b8-11ea-832a-68f72850e714    Todo    tag2    Content 2
e0da29df-80b8-11ea-8f7f-68f72850e714    Todo    tag3    Content 3

Outdated tasks: 0
```

Let's set specific task as `Done` and check the list again.

```sh
❯ ./efishery-cli-app done -i d63c76b0-80b8-11ea-983a-68f72850e714
❯ ./efishery-cli-app list

docID                                   Status  Tags    Content
d63c76b0-80b8-11ea-983a-68f72850e714    Done    tag1    Content 1
dd6badac-80b8-11ea-832a-68f72850e714    Todo    tag2    Content 2
e0da29df-80b8-11ea-8f7f-68f72850e714    Todo    tag3    Content 3

Outdated tasks: 0
```

Let's try to delete a specific task.

```sh
❯ ./efishery-cli-app delete -i dd6badac-80b8-11ea-832a-68f72850e714
❯ ./efishery-cli-app list

docID                                   Status  Tags    Content
d63c76b0-80b8-11ea-983a-68f72850e714    Done    tag1    Content 1
e0da29df-80b8-11ea-8f7f-68f72850e714    Todo    tag3    Content 3

Outdated tasks: 0
```

Let's try to update specific task.

```sh
❯ ./efishery-cli-app edit -c "Content 4" -t tag3,tag4 -i e0da29df-80b8-11ea-8f7f-68f72850e714
❯ ./efishery-cli-app list

docID                                   Status  Tags            Content
d63c76b0-80b8-11ea-983a-68f72850e714    Done    tag1            Content 1
e0da29df-80b8-11ea-8f7f-68f72850e714    Todo    tag3,tag4       Content 4

Outdated tasks: 0
```

To check the offline-first feature, please disconnect your internet and check the list again.

```sh
❯ ./efishery-cli-app list

docID                                   Status  Tags            Content
d63c76b0-80b8-11ea-983a-68f72850e714    Done    tag1            Content 1
e0da29df-80b8-11ea-8f7f-68f72850e714    Todo    tag3,tag4       Content 4

Outdated tasks: 0
```

Add a task and edit another task.

```sh
❯ ./efishery-cli-app add -c "Content 5" -t tag5
❯ ./efishery-cli-app done -i e0da29df-80b8-11ea-8f7f-68f72850e714
❯ ./efishery-cli-app list

> No internet connection, using local data

docID                                   Status  Tags            Content
d63c76b0-80b8-11ea-983a-68f72850e714    Done    tag1            Content 1
e0da29df-80b8-11ea-8f7f-68f72850e714    Done    tag3,tag4       Content 4
636be3cb-80ba-11ea-b95d-68f72850e714    Todo    tag5            Content 5

Outdated tasks: 0
```

Now, connect your internet connection and check the list. The list will show you a different tasks since it automatically use remote database.

```sh
❯ ./efishery-cli-app list

docID                                   Status  Tags            Content
d63c76b0-80b8-11ea-983a-68f72850e714    Done    tag1            Content 1
e0da29df-80b8-11ea-8f7f-68f72850e714    Todo    tag3,tag4       Content 4

Outdated tasks: 1
```

Let's sync the tasks to update the remote database.

```sh
❯ ./efishery-cli-app sync
❯ ./efishery-cli-app list

docID                                   Status  Tags            Content
636be3cb-80ba-11ea-b95d-68f72850e714    Todo    tag5            Content 5
d63c76b0-80b8-11ea-983a-68f72850e714    Done    tag1            Content 1
e0da29df-80b8-11ea-8f7f-68f72850e714    Todo    tag3,tag4       Content 4

Outdated tasks: 0
```
