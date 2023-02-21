use crate::schema::users;
use diesel::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Insertable, Serialize, Deserialize)]
#[table_name = "users"]
pub struct NewUser {
    pub username: String,
    pub token: String,
    pub wins: i32,
    pub losses: i32,
}

#[derive(Debug, Serialize, Queryable, AsChangeset, Insertable)]
pub struct User {
    pub id: i32,
    pub username: String,
    pub token: String,
    pub wins: i32,
    pub losses: i32,
}
