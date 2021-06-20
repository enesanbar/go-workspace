from datetime import date
from datetime import time
from datetime import datetime


def main():
    # DATE OBJECTS
    # Get today's date from the simple today() method from the date class
    today = date.today()
    print("Today's date is ", today)

    # print out the date's individual components
    print("Date Components: ", today.day, today.month, today.year)

    # retrieve today's weekday (0=Monday, 6=Sunday)
    print("Today's Weekday #: ", today.weekday())
    days = ["monday", "tuesday", "wednesday",
            "thursday", "friday", "saturday", "sunday"]
    print("Which is a " + days[today.weekday()])

    # DATETIME OBJECTS
    # Get today's date from the datetime class
    now = datetime.now()
    print("The current date and time is ", now)

    # Get the current time`
    t = datetime.time(datetime.now())
    print("The current time is ", t)

    #### Date Formatting ####
    # %y/%Y - Year, %a/%A - weekday, %b/%B - month, %d - day of month
    print(now.strftime("The current year is: %Y"))  # full year with century
    # abbreviated day, num, full month, abbreviated year
    print(now.strftime("%a, %d %B, %y"))

    # %c - locale's date and time, %x - locale's date, %X - locale's time
    print(now.strftime("Locale date and time: %c"))
    print(now.strftime("Locale date: %x"))
    print(now.strftime("Locale time: %X"))

    #### Time Formatting ####
    # %I/%H - 12/24 Hour, %M - minute, %S - second, %p - locale's AM/PM
    # 12-Hour:Minute:Second:AM
    print(now.strftime("Current time: %I:%M:%S %p"))
    print(now.strftime("24-hour time: %H:%M"))  # 24-Hour:Minute


if __name__ == "__main__":
    main()
