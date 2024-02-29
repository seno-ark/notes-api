CREATE TABLE "notes" (
  "id" VARCHAR(26) PRIMARY KEY,
  "title" VARCHAR(100) NOT NULL,
  "content" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);