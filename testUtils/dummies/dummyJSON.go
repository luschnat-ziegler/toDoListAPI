package dummies

var DummyListValidAsJSON = `{"id":"000000000000000000000000","name":"Dummy List Name","description":null,"tasks":[{"id":"","name":"Dummy Task 1","description":null},{"id":"","name":"Dummy Task 2","description":null}]}`
var DummyListValidWithIdsAsJson = `{"id":"601be448b9b5e15374b1e842","name":"Dummy List Name","description":null,"tasks":[{"id":"1234","name":"Dummy Task 1","description":null},{"id":"3245","name":"Dummy Task 2","description":null}]}`
var DummyRequestInvalidJSON = `{id":"601be448b9b5e15374b1e842","name":"Dummy List Name","description":null,"tasks":[{"id":"1234","name":"Dummy Task 1","description":null},{"id":"3245","name":"Dummy Task 2","description":null}]}`
var DummyInternalErrorAsJSON = `{"message":"internal error"}`
var DummyBadRequestErrorAsJSON = `{"message":"Body parsing error"}`
var DummyValidSaveListRequestAsJSON = `{"name":"Dummy List Name", "description":null, "tasks":[{"name":"Dummy Task 1","description":null},{"name":"Dummy Task 2","description":null}]}`
var DummyInvalidSaveListRequestAsJSON = `{"name":"", "description":null, "tasks":[{"name":"Dummy Task 1","description":null},{"name":"Dummy Task 2","description":null}]}`
var DummyValidationErrorAsJSON = `{"invalid_fields":{"name":"required"}}`
