# Eventurely Backend

Eventurely is an application that allows users to create and manage events. It is created using Go, Connect RPC, and Flutter. Created as a learning exercise whilst (loosely) following the book "Let's Go" by Alex Edwards.

## Dependencies

- Go v1.21
- Air v1.49 (optional, for hot reloading)
- Buf v1.29 (optional, for generating Protobuf files)

## Notes
- This project is a work in progress and is not yet ready for production use.
- The project leverages the `buf` CLI for generating the Protobuf files for the gRPC server. This can be done using `buf generate api/proto` from the root of the project. This will place the generated files in the `gen` directory.

## MVP Stories

### User Stories

As a person using the app for the first time, I want to...

- [ ] continue as a guest, providing minimal personal data.
- [ ] create an account, becoming a 'registered' user.

As a guest, I want to...

- [ ] view a list of my upcoming events (invited to).
- [ ] RSVP to a future event I am invited to.
- [ ] receive notifications about my upcoming events.
- [ ] view past events that I owned or was invited to.
- [ ] create an account, becoming a registered user, without losing any of my event data.

As a registered user, I want to...

- [ ] log in to my account.
- [ ] log out of my account.
- [ ] make changes to my account details.
- [ ] create a new event.
- [ ] view a list of my upcoming events (owned and invited to).
- [ ] edit all the details of my future event.
  - [ ] choose whether my guests are notified about the change/s to event details.
- [ ] delete my future event.
- [ ] invite anyone to my future event using their email, phone number, or by sharing a unique link.
- [ ] see the RSVP status (Invited, Yes, No, Maybe) of each invitee of my event.
- [ ] RSVP to a future event I am invited to. 
- [ ] receive notifications about my upcoming events.
- [ ] view past events that I owned or was invited to.

### Engineer Stories

As an engineer, I want to...

#### Mobile App

- [ ] hide/change certain functionality if the user is a guest, in accordance with the User Stories.
- [ ] implement a form for event creation.
- [ ] implement a form for event editing. 
- [ ] implement a delete event option, with a confirmation message.
- [ ] implement a list of events belonging to a user, sorted by date.
- [ ] integrate an email service for sending event invitations and updates.
- [ ] integrate a phone service for sending event invitations and updates.
- [ ] create a notification system to alert users about upcoming events, detail changes, RSVPs, etc.
- [ ] restrict user actions on past events.

#### Web API

##### Endpoints

The following endpoints are divided by the views/screens they are most associated with. This does not mean they are restricted to those views.

**Welcome/Login/Signup:**

Coming soon...

**View/Create/Edit:**

- [x] **EventService.Get**: get an event.
- [x] **EventService.Create**: create a new event.
- [ ] **EventService.Update**: edit an event that the user owns.
- [ ] **EventService.Delete**: delete an event that the user owns.
- [ ] invite a user to an event.
  - ! This (and similar invite endpoints) will likely not be its own endpoint, but rather a part of the event creation/editing process.

**Pending Invites:**

_NOTE: This will default to invites that have not been responded to, but filtering will be available._

- [ ] **EventService.ListUpcomingInvited**: get the event & invitation for each upcoming event that a user has been invited to.
  - This should take a filter, and default to only showing events the user has not responded to.
- [ ] update the RSVP status of an invitation.

**Upcoming:**

- [ ] **EventService.ListUpcomingInvited**: get the event & invitation for each upcoming event that a user has been invited to.
  - This should take a filter, and default to only showing events the user had responded "Yes" to.

**Archive:**

- [ ] **EventService.ListPast**: get the event & invitation for each past event that a user has been invited to or owned.
  - This should take a filter, and default to showing all events.

**Organising:**

- [ ] **EventService.ListUpcomingOwned**: get the event & invitation for each event that a user owns.
  - This should take a filter, and default to showing all events.

**Connections:**

Coming soon...

**Settings:**

Coming soon...