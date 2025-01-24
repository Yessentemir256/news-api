-- news table
CREATE TABLE "news" (
  "id" bigserial PRIMARY KEY,
  "title" text NOT NULL,
  "content" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

-- news_categories table
CREATE TABLE "news_categories" (
  "news_id" bigint NOT NULL,
  "category_id" bigint NOT NULL,
  PRIMARY KEY ("news_id", "category_id")
);

-- Indexes for news table
CREATE INDEX ON "news" ("title");
CREATE INDEX ON "news" ("created_at");

-- Indexes for news_categories table
CREATE INDEX ON "news_categories" ("news_id");
CREATE INDEX ON "news_categories" ("category_id");

-- Foreign key constraints
ALTER TABLE "news_categories" 
ADD FOREIGN KEY ("news_id") REFERENCES "news" ("id") ON DELETE CASCADE;

-- Optional comments for better schema documentation
COMMENT ON COLUMN "news"."title" IS 'Title of the news item';
COMMENT ON COLUMN "news"."content" IS 'Content of the news item';
COMMENT ON TABLE "news_categories" IS 'Linking table for news and categories';
