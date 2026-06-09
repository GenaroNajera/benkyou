# benkyou (べんくょう)

This project helps with studying the kyouiku kanji, the characters Japanese students learn in elementary school.
It currently only includes the characters learned in first grade.
Clicking the kanji opens its Wiktionary page for more details.

# Run

To create the database, go to the database folder then, in the terminal, type
```sql
sqlite3 benkyou.db
.read create_insert.sql
```

Start the server with
```go
go run main.go
```
and go to localhost:8080 in the browser.

![home](/images/home.jpg)

Click 'Start Studying!' to start studying!

![study](/images/study_1.jpg)
![study](/images/study_2.jpg)

The information is styled as a flashcard with only the kanji revealed. The meaning and readings are hidden and are revealed when clicking the 'Reveal' button.
The 'Next' and 'Prev' buttons navigate to the other kanji.