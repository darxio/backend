CREATE EXTENSION pg_trgm;
create index on products_extended using gin (lower(name) gin_trgm_ops, lower(category_url) gin_trgm_ops);