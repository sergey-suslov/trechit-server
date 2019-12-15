cd ../../timer && dotenv -e .env psql "$DATABASE_NAME" -- -c "select 'drop table if exists "' || tablename || '" cascade;'
  from pg_tables
 where schemaname = 'public';"


