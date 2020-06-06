# scheduler

TBD

## Development

TBD

Building:
```sh
make build
```

Starting:
```sh
make start
```

Testing:
```sh
make go-test
```

Running:
```sh
make run
```

Cleaning:
```sh
make clean
```

Executing all commands:
```sh
make all
```

## Usage

TBD

### Files format

#### Groups

| Name | Type | Special format | Description |
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
| group | General | - | Group name |
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
