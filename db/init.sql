
create table if not exists hook_event (
  event_id SERIAL PRIMARY KEY,
  event_type TEXT,
  event_data TEXT,
  status TEXT DEFAULT 'pending',

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
