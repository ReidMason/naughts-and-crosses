mod controllers;
mod database_models;
mod schema;
mod services;

use std::io::Result;

use actix_web::{middleware::Logger, web, App, HttpServer};
use controllers::user_controller::{create_user, get_user, get_users};
use diesel::{r2d2::ConnectionManager, SqliteConnection};

type DbPool = r2d2::Pool<ConnectionManager<SqliteConnection>>;

#[actix_web::main]
async fn main() -> Result<()> {
    // Setup the logger
    std::env::set_var("RUST_LOG", "debug");
    std::env::set_var("RUST_BACKTRACE", "1");
    env_logger::init();

    dotenv::dotenv().ok();
    let conn_spec = std::env::var("DATABASE_URL").expect("Failed to find database url");
    let manager = ConnectionManager::<SqliteConnection>::new(conn_spec);
    let pool = r2d2::Pool::builder()
        .build(manager)
        .expect("Failed to create database pool");

    // Services
    let user_service = services::user_service::UserService {};

    // Initialise the web server
    HttpServer::new(move || {
        // pass in the logger
        let logger = Logger::default();
        App::new()
            .app_data(web::Data::new(pool.clone()))
            .app_data(web::Data::new(user_service.clone()))
            .wrap(logger)
            .service(get_user)
            .service(create_user)
            .service(get_users)
    })
    .bind(("0.0.0.0", 3000))?
    .run()
    .await
}
