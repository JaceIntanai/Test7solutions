# Test7solutions

## Question 1
### Run read json
```bash
go run question1.go
```
### Explain

>The working principle is to calculate by comparing left and right from the bottom node to the top. You will start from the second row from the bottom and then look at the nodes below to see which one is more, then use that number and combine it with itself. Then move up the row, and finally, when you reach the top node, it will choose the greatest left or right and combine them with itself to get the greatest distance total.


## Question 2
### Run to CLI form keyboard(encoded)
```bash
go run question2.go
```
### Explain

>The working principle starts by receiving encoded values ​​and then checking the characters from front to back. The first one after that is assumed to be the lowest possible value and then moves through the letters. If there is a chance that the number will be negative, it increases the value of the previous number to make the letter can be used after getting the result, then fill in the value according to the number specified. If it is a = sign, it will add the number to match the side. If not, it will add 0 to complete the number.


## Question 3
### Run Json API
```bash
go run question3.go
```
### Explain
```
POST Api : http://localhost:8080/beef/summary?text=...
```
>The working principle of this api is to receive values ​​from a variable named text and then convert them into fields list in order to group by the names with matching names, make the names are all lowercase letters and remove unnecessary characters such as '.' or ',' and put them into matching groups for counting. After finishing, convert the values ​​to json to send the response back.
Example: text = (value form https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text)
```
Result: {
    "beef": {
        "ad": 47,
        "adipisicing": 40,
        "alcatra": 30,
        "aliqua": 37,
        "aliquip": 33,
        "andouille": 46,
        "anim": 38,
        "aute": 27,
        "bacon": 41,
        "ball": 34,
        "beef": 113,
        "belly": 39,
        "biltong": 49,
        "boudin": 49,
        "bresaola": 43,
        "brisket": 43,
        "buffalo": 41,
        "burgdoggen": 33,
        "capicola": 42,
        "chicken": 42,
        "chislic": 38,
        "chop": 51,
        "chuck": 32,
        "cillum": 42,
        "commodo": 37,
        "consectetur": 39,
        "consequat": 38,
        "corned": 37,
        "cow": 50,
        "culpa": 42,
        "cupidatat": 46,
        "cupim": 29,
        "deserunt": 38,
        "do": 35,
        "dolor": 41,
        "dolore": 79,
        "doner": 44,
        "drumstick": 35,
        "duis": 41,
        "ea": 44,
        "eiusmod": 36,
        "elit": 40,
        "enim": 53,
        "esse": 43,
        "est": 41,
        "et": 37,
        "eu": 37,
        "ex": 35,
        "excepteur": 43,
        "exercitation": 42,
        "fatback": 35,
        "filet": 37,
        "flank": 30,
        "frankfurter": 25,
        "fugiat": 38,
        "ground": 46,
        "ham": 89,
        "hamburger": 35,
        "hock": 49,
        "id": 36,
        "in": 97,
        "incididunt": 34,
        "ipsum": 36,
        "irure": 28,
        "jerky": 46,
        "jowl": 31,
        "kevin": 43,
        "kielbasa": 31,
        "labore": 44,
        "laboris": 50,
        "laborum": 31,
        "landjaeger": 44,
        "leberkas": 39,
        "loin": 74,
        "lorem": 28,
        "magna": 33,
        "meatball": 40,
        "meatloaf": 42,
        "mignon": 37,
        "minim": 43,
        "mollit": 44,
        "nisi": 42,
        "non": 35,
        "nostrud": 35,
        "nulla": 51,
        "occaecat": 49,
        "officia": 30,
        "pancetta": 36,
        "pariatur": 45,
        "pastrami": 49,
        "picanha": 36,
        "pig": 32,
        "porchetta": 44,
        "pork": 157,
        "proident": 36,
        "prosciutto": 34,
        "qui": 37,
        "quis": 43,
        "reprehenderit": 41,
        "ribeye": 32,
        "ribs": 108,
        "round": 46,
        "rump": 42,
        "salami": 35,
        "sausage": 39,
        "sed": 32,
        "shank": 40,
        "shankle": 48,
        "short": 63,
        "shoulder": 45,
        "sint": 42,
        "sirloin": 43,
        "spare": 50,
        "steak": 31,
        "strip": 31,
        "sunt": 42,
        "swine": 34,
        "t-bone": 40,
        "tail": 51,
        "tempor": 30,
        "tenderloin": 42,
        "tip": 34,
        "tongue": 33,
        "tri-tip": 34,
        "turducken": 34,
        "turkey": 28,
        "ullamco": 30,
        "ut": 125,
        "velit": 44,
        "veniam": 43,
        "venison": 42,
        "voluptate": 46
    }
}
```
