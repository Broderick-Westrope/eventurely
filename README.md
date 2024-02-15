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
- [ ] filter my Pending Invites, Upcoming, Archive, and Organising lists by:
  - Events that occur within a given time frame (startsAfter & endsBefore)
  - Events with only a given privacy setting
  - Invited events whose invitation has a given RSVP status (not applicable to Organising list) 

_NOTE: Guest functionality will be added at a later date._

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

- [x] **GetEvent**: get an event.
- [x] **CreateEvent**: create a new event.
- [ ] **UpdateEvent**: edit an event that the user owns.
- [ ] **DeleteEvent**: delete an event that the user owns.
- [ ] invite a user to an event.
  - ! This (and similar invite endpoints) will likely not be its own endpoint, but rather a part of the event creation/editing process.

**Pending Invites:**

_NOTE: This will default to invites that have not been responded to, but filtering will be available._

- [x] **ListUpcomingInvitedEvents**: get the event & invitation for each upcoming event that a user has been invited to.
  - This should take a filter on the frontend, and default to only showing events the user has not responded to.
- [x] update the RSVP status of an invitation.
  - [x] When the user resets the response to "Sent" the respondedAt time should be set to null. If the response status is set to anything else, the respondedAt time should be updated to the current time.

**Upcoming:**

- [x] **ListUpcomingInvitedEvents**: get the event & invitation for each upcoming event that a user has been invited to.
  - [ ] This should take a filter on the frontend, and default to only showing events the user had responded "Yes" to.

**Archive:**

- [x] **ListPastEvents**: get the event & invitation for each past event that a user has been invited to or owned.
  - [ ] This should take a filter on the frontend, and default to showing all events.

**Organising:**

- [x] **ListUpcomingOwnedEvents**: get the event & invitation for each event that a user owns.
  - [ ] This should take a filter on the frontend, and default to showing all events.

**Connections:**

Coming soon...

**Settings:**

Coming soon...

### Future Additions

Following is a list of low-priority features that could be added at a later date, not including those mentioned elsewhere in this document:
- [ ] collaborative events: allow users to add co-owners who also have the same rights to editing the event. Should co-owners be able to kick others or only leave willingly?
- [ ] several emails per account, each with their own verification check