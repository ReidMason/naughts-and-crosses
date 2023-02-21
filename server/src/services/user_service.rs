use crate::database_models::users::{NewUser, User};
use diesel::prelude::*;
use diesel::{OptionalExtension, QueryDsl, RunQueryDsl, SqliteConnection};
use std::io::{Error, ErrorKind};
use uuid::Uuid;

#[derive(Clone)]
pub struct UserService;

type DbError = Box<dyn std::error::Error + Send + Sync>;

pub trait UserServiceInterface {
    fn get_users(&self, database_connection: &mut SqliteConnection) -> Result<Vec<User>, DbError>;
    fn get_user(
        &self,
        database_connection: &mut SqliteConnection,
        user_id: i32,
    ) -> Result<Option<User>, DbError>;
    fn create_user(
        &self,
        database_connection: &mut SqliteConnection,
        username: String,
    ) -> Result<Option<NewUser>, DbError>;
}

impl UserServiceInterface for UserService {
    fn get_users(&self, database_connection: &mut SqliteConnection) -> Result<Vec<User>, DbError> {
        use crate::schema::users::dsl::*;

        let found_users = users.load::<User>(database_connection)?;

        Ok(found_users)
    }

    fn get_user(
        &self,
        database_connection: &mut SqliteConnection,
        user_id: i32,
    ) -> Result<Option<User>, DbError> {
        use crate::schema::users::dsl::*;

        let found_user = users
            .filter(id.eq(user_id))
            .first::<User>(database_connection)
            .optional()?;
        Ok(found_user)
    }

    fn create_user(
        &self,
        database_connection: &mut SqliteConnection,
        new_username: String,
    ) -> Result<Option<NewUser>, DbError> {
        use crate::schema::users::dsl::*;

        let existing_user = users
            .filter(username.eq(new_username.clone()))
            .first::<User>(database_connection)
            .optional()?;

        if existing_user.is_some() {
            // errors can be created from strings
            let new_error = Error::new(
                ErrorKind::AlreadyExists,
                "A user with that username already exists",
            );
            return Err(Box::new(new_error));
        }

        let new_user = NewUser {
            username: new_username,
            token: Uuid::new_v4().to_string(),
            wins: 0,
            losses: 0,
        };
        diesel::insert_into(users)
            .values(&new_user)
            .execute(database_connection)?;

        Ok(Some(new_user))
    }
}
