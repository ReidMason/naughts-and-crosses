use crate::{
    schema::users::dsl::*,
    services::user_service::{UserService, UserServiceInterface},
    DbPool,
};
use actix_web::{
    error::ResponseError,
    get,
    http::{header::ContentType, StatusCode},
    post, put,
    web::Data,
    web::Path,
    web::{self, Json},
    Error, HttpResponse,
};
use derive_more::Display;
use diesel::{
    prelude::*,
    r2d2::{self, ConnectionManager},
};
use diesel::{RunQueryDsl, SqliteConnection};
use serde::{Deserialize, Serialize};
use uuid::Uuid;

use crate::database_models::users::{self, NewUser};

#[derive(Serialize, Deserialize)]
pub struct UserDTO {
    username: String,
    wins: i32,
    losses: i32,
}

#[derive(Serialize, Deserialize)]
pub struct NewUserDTO {
    username: String,
}

#[derive(Serialize, Deserialize)]
pub struct UserIdentifier {
    user_id: i32,
}

#[derive(Serialize, Deserialize)]
pub struct UpdateUserDTO {
    user_id: String,
    token: String,
}

#[get("/user/{user_id}")]
pub async fn get_user(
    db_pool: web::Data<DbPool>,
    user_service: web::Data<UserService>,
    user_identifier: Path<UserIdentifier>,
) -> Result<HttpResponse, Error> {
    let user_id = user_identifier.into_inner().user_id;
    let found_user = web::block(move || {
        let mut conn = db_pool.get()?;
        user_service.get_user(&mut conn, user_id)
    })
    .await?
    .map_err(actix_web::error::ErrorInternalServerError)?;

    match found_user {
        Some(found_user) => Ok(HttpResponse::Ok().json(found_user)),
        None => Ok(HttpResponse::NotFound().body(format!("No user found with uid: {user_id}"))),
    }
}

// #[put("/user/{user_id}")]
// pub async fn update_user(
//     pool: web::Data<DbPool>,
//     user_identifier: Path<UserIdentifier>,
//     body: Json<UpdateUserDTO>,
// ) -> Json<String> {
//     let updated_user = body.into_inner();

//     let mut connection = pool
//         .get()
//         .expect("Failed to get database connection from pool");

//     let insert = diesel::insert_into(users)
//         .values(&new_user)
//         .execute(&mut connection);

//     Json(user_identifier.into_inner().user_id)
// }

#[get("/user")]
async fn get_users(
    db_pool: web::Data<DbPool>,
    user_service: web::Data<UserService>,
) -> Result<HttpResponse, Error> {
    let found_users = web::block(move || {
        let mut conn = db_pool.get()?;
        user_service.get_users(&mut conn)
    })
    .await?
    .map_err(actix_web::error::ErrorInternalServerError)?;

    Ok(HttpResponse::Ok().json(found_users))
}

#[post("/user")]
pub async fn create_user(
    db_pool: web::Data<DbPool>,
    user_service: web::Data<UserService>,
    body: Json<NewUserDTO>,
) -> Result<HttpResponse, Error> {
    let new_username = body.into_inner().username;

    let new_user = web::block(move || {
        let mut conn = db_pool.get()?;
        user_service.create_user(&mut conn, new_username)
    })
    .await?
    .map_err(actix_web::error::ErrorInternalServerError)?;

    Ok(HttpResponse::Ok().json(new_user))
}
