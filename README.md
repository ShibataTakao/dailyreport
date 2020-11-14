# daily-report

[![CircleCI](https://circleci.com/gh/ShibataTakao/daily-report.svg?style=shield)](https://circleci.com/gh/ShibataTakao/daily-report)

## Usage

### Create daily-report
You can create daily-report by following command.

```
$ dailyrepot create
今日の日報 /home/ubuntu/note/20200229.md を作成しました。
```

Daily-report will be created at `${DAILYREPORT_PATH}/yyyymmdd.md` .

You can change daily-report template by rewriting [template.go](./internal/dailyreport/template.go) if you want.

### Validate daily-report
You can show worktime summary of today's daily-report to validate it.

```
$ dailyreport validate
業務時間 = 7.00h
今日のタスク（予定） = 0.00h
今日のタスク（実績） = 0.00h
```

### Aggregate daily-report
You can show aggregated tasks for daily-reports in specified date range.

```
$ dailyreport aggregate --from 20201001 --to 20201031
```

## Install
```
make install
```
