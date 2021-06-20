import urllib.request
import json


def print_results(data):
    # Use the json module to load the string data into a dictionary
    json_data = json.loads(data)

    # now we can access the contents of the JSON like any other Python object
    if "title" in json_data["metadata"]:
        print(json_data["metadata"]["title"])

    # output the number of events, plus the magnitude and each event name
    count = json_data["metadata"]["count"]
    print(str(count) + " events recorded")

    # for each event, print the place where it occurred
    for i in json_data["features"]:
        print(i["properties"]["place"])
    print("--------------\n")

    # print the events that only have a magnitude greater than 4
    for i in json_data["features"]:
        if i["properties"]["mag"] >= 4.0:
            print("%2.1f" % i["properties"]["mag"], i["properties"]["place"])
    print("--------------\n")

    # print only the events where at least 1 person reported feeling something
    print("\n\nEvents that were felt:")
    for feature in json_data["features"]:
        felt_reports = feature["properties"]["felt"]
        if felt_reports:
            print("%2.1f" % feature["properties"]["mag"], feature["properties"]
                  ["place"], " reported " + str(felt_reports) + " times")
            print(
                f"{feature['properties']['mag']} {feature['properties']['place']} reported {felt_reports} times")


def main():
    # This feed lists all earthquakes for the last day larger than Mag 2.5
    url = "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/2.5_day.geojson"

    # Open the URL and read the data
    response = urllib.request.urlopen(url)
    print("result code: " + str(response.getcode()))
    if (response.getcode() == 200):
        data = response.read().decode("utf-8")
        print_results(data)
    else:
        print("Received an error from server, cannot retrieve results " +
              str(response.getcode()))


if __name__ == "__main__":
    main()
