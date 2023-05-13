# worklog

[![Test](https://github.com/ShibataTakao/worklog/actions/workflows/test.yaml/badge.svg)](https://github.com/ShibataTakao/worklog/actions/workflows/test.yaml)

CLI tool to control worklog.

## Usage

### Show work time
```
Show work time in worklog.

Usage:
  worklog show worktime [flags]

Flags:
      --daily-report-dir string        Directory where daily report file exists. [$WL_DAILY_REPORT_DIR]
  -e, --daily-report-end-at string     End of daily report date range. [$WL_DAILY_REPORT_END_AT]
  -s, --daily-report-start-at string   Start of daily report date range. [$WL_DAILY_REPORT_START_AT]
      --filter-by-project string       Show only tasks which project name is this. [$WL_FILTER_BY_PROJECT]
  -h, --help                           help for worktime
```

### Show tasks
```
Show tasks in worklog.

Usage:
  worklog show tasks [flags]

Flags:
      --backlog-api-key string             Backlog API key. [$WL_BACKLOG_API_KEY]
      --backlog-project-overwrite string   Project name of issues in backlog will be overwritten by this. [$WL_BACKLOG_PROJECT_OVERWRITE]
      --backlog-query string               Query to get issues from backlog. [$WL_BACKLOG_QUERY]
      --backlog-url string                 Backlog URL. [$WL_BACKLOG_URL]
      --daily-report-dir string            Directory where daily report file exists. [$WL_DAILY_REPORT_DIR]
  -e, --daily-report-end-at string         End of daily report date range. [$WL_DAILY_REPORT_END_AT]
  -s, --daily-report-start-at string       Start of daily report date range. [$WL_DAILY_REPORT_START_AT]
      --filter-by-project string           Show only tasks which project name is this. [$WL_FILTER_BY_PROJECT]
  -h, --help                               help for tasks
```

## Install
```
go install github.com/ShibataTakao/worklog@master
```
