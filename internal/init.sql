CREATE TABLE IF NOT EXISTS "isbns" (
  "id" SERIAL PRIMARY KEY,
  "v10" VARCHAR(20),
  "v13" VARCHAR(20),
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE IF NOT EXISTS "series" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "ongoing" BOOLEAN DEFAULT true,
  "total_books" INTEGER,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE IF NOT EXISTS "books" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR(500) NOT NULL,
  "series_id" INTEGER,
  "book_in_series" INTEGER,
  "publisher" VARCHAR(255),
  "publish_date" DATE,
  "cover_url" TEXT,
  "media_url" TEXT,
  "isbn_id" INTEGER,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE IF NOT EXISTS "genres" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE IF NOT EXISTS "books_genres" (
  "book_id" INTEGER,
  "genre_id" INTEGER,
  PRIMARY KEY ("book_id", "genre_id")
);

CREATE TABLE IF NOT EXISTS "authors" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE IF NOT EXISTS "books_authors" (
  "book_id" INTEGER,
  "author_id" INTEGER,
  PRIMARY KEY ("book_id", "author_id")
);

ALTER TABLE "books" ADD FOREIGN KEY ("series_id") REFERENCES "series" ("id");

ALTER TABLE "books" ADD FOREIGN KEY ("isbn_id") REFERENCES "isbns" ("id");

ALTER TABLE "books_genres" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "books_genres" ADD FOREIGN KEY ("genre_id") REFERENCES "genres" ("id");

ALTER TABLE "books_authors" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "books_authors" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");
