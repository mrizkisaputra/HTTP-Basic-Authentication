package utils

const QueryContacts = "SELECT id, name, phone_number, created_at FROM contacts"
const QueryByIdContact = "SELECT * FROM contacts WHERE id = ?"
