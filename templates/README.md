## Groupie Tracker

Groupie Tracker is a web application built to display and manipulate data from a given API containing information about artists, their concert locations, dates, and relationships between these elements. The application provides a user-friendly interface where users can browse artist information, view concert locations and dates, and interact with server-side data in real-time using client-server events.

### Project Overview

Groupie Tracker displays detailed information about music artists and their concerts by consuming an API. The API provides information in four parts:

1. Artists - Information about bands and artists, including name, image, activity year, first album, and band members.
2. Locations - Details about the upcoming or past concert locations.
3. Dates - Concert dates.
4. Relations - Links artists, locations, and concert dates.

The application aims to provide rich visualizations of this data using cards, lists, and tables, making it easy for users to explore information interactively.
### Features

- Artist Information: Displays artists' details, including band members, start date, and first album.
- Concert Locations & Dates: Shows upcoming and past concert locations and dates.
- Client-Server Interaction: Implements a feature that triggers a server request and updates information dynamically.
- Error Handling: Ensures that the site is robust and handles any potential errors gracefully.
- Mobile-Responsive Design: The layout adjusts to different screen sizes for a better user experience on mobile devices.

### Technologies Used

- Backend: Go (Golang) for server-side logic and API interaction.
- Frontend: HTML, CSS for building a dynamic and responsive user interface.
- Template Engine: Go's html/template for rendering HTML pages.
- HTTP Server: Go's net/http package to serve pages and handle requests.
- Client-Server Communication: AJAX or Fetch API to communicate between the client and server without reloading the page.
- Unit Testing: Go's testing framework for ensuring code correctness.

### Setup Instructions
#### Prerequisites

- Go 1.20+ installed on your machine.
- Internet connection to fetch the API data.

#### Steps to Run

1. Clone the repository:

```bash
git clone https://learn.zone01kisumu.ke/git/joseowino/groupie-tracker
```

2. Navigate into the project directory:

```bash
cd groupie-tracker
```

3. Install dependencies (if any).

4. Run the application:
```bash
go run main.go
```

5. Open your browser and navigate to http://localhost:8080.

### API Structure

The application interacts with the following parts of the API:
#### Artists

- Endpoint: /api/artists
- Details: Provides artist details like name, members, first album date, and an image.

#### Locations

- Endpoint: /api/locations
- Details: Contains a list of cities where the artist has performed or will perform.

#### Dates

- Endpoint: /api/dates
- Details: Lists the dates of the artist's concerts.

#### Relations

- Endpoint: /api/relations
- Details: Links between artists, locations, and dates.

### Error Handling

- 404 Page: A custom "Not Found" page is shown if a user navigates to an unknown URL.
- 500 Internal Server Error: A custom page is displayed when a server-side error occurs.
- Validation: All inputs are validated on the server and client to avoid malformed requests.

### Testing

Unit tests ensure the correctness of the backend code. Tests are implemented using Go's testing package.

To run the tests:

```bash
go test 
```

### Future Improvements

- Advanced Data Visualization: Adding charts to visualize concert dates and locations.
- Search Feature: Implementing a search functionality to quickly find artists or concerts.
- User Authentication: Allowing users to create accounts and follow their favorite artists.