// @generated automatically by Diesel CLI.

diesel::table! {
    users (id) {
        id -> Integer,
        username -> Text,
        token -> Text,
        wins -> Integer,
        losses -> Integer,
    }
}
