# scheduler

Scheduler is a tool for creating students' schedule based on their preferences. It receives four paths to files as arguments. The files are described in the Usage section.

## Development

Please note that the makefile won't work on Windows OS.

[Go language](https://golang.org/doc/install)

Building:
```sh
make build
```

Starting:
```sh
make start groups=./path/to/groups.xlsx students=./path/to/students/directory priority=./path/to/priority_students.xlsx result=./path/to/results/directory
```

Testing:
```sh
make go-test
```

Running:
```sh
make run groups=./path/to/groups.xlsx students=./path/to/students/directory priority=./path/to/priority_students.xlsx result=./path/to/results/directory
```

Cleaning:
```sh
make clean
```

Executing all commands:
```sh
make all groups=./path/to/groups.xlsx students=./path/to/students/directory priority=./path/to/priority_students.xlsx result=./path/to/results/directory
```

Arguments:

| Argument | Default Value | Description |
| -------- | ------------- | ----------- |
| groups | ./example/groups.xlsx | Path to a file which contains groups |
| students | ./example/students | Path to a directory which contains students preferences |
| priority | ./example/priority_students.xlsx | Path to a file which contains list of priority students |
| result | ./example/result | Path to a directory where the results will be saved |

## Usage

Please note that the `example` directory contains sample files and directories.

Linux / OSX:

```sh
./main -groups=./path/to/groups.xlsx -students=./path/to/students/directory -priority=./path/to/priority_students.xlsx -result=./path/to/results/directory
```

Windows:

```sh
.\main.exe -groups=".\path\to\groups.xlsx" -students=".\path\to\students\directory" -priority=".\path\to\priority_students.xlsx" -result=".\path\to\results\directory"
```

### Files structures

#### Groups

| Name | Type | Format | Description |
| ---- | ---- | ----- | ----------- |
| name | General | - | Subject name |
| type | General | Lecture &#124; Laboratory &#124; Class | Type of a class |
| teacher | General | - | Teacher name |
| weekday | General | Monday &#124; Tuesday &#124; Wednesday &#124; Thursday &#124; Friday | Day on which classes are held |
| start time | Text | hour:minutes, e.g. 15:04 | Start time of a group |
| end time | Text | hour:minutes, e.g. 15:04 | End time of a group |
| place | General | - | Place where classes are held |
| start date | Date | day/month/year, e.g. 02/01/2006 | Start date of a group |
| frequency | Number | - | How often class is held: 1 - 1/1 week, 2 - 1/2 weeks, etc. |
| group | General | - | Group name, use Lecture if group type is set to Lecture |
| capacity | Number | - | Maximum number of students per group |   

#### Student

| Name | Type | Description |
| ---- | ---- | ----------- |
| name | General | Subject name |
| group | General | Group name |
| priority | Number | How much a group is important for a student |

Please note that the `priorities` within one subject must be consecutive and start from 1 (the most important group). They can be repeated.

Please note that the `file name` will be parsed as a `student's name`.

#### Priority Students

| Name | Type | Description |
| ---- | ---- | ----------- |
| name | General | Name of a priority student |
