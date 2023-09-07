CREATE TABLE todo (
		id TEXT not null primary key, 
		content TEXT
	); 
CREATE UNIQUE INDEX idx_data_id ON todo(id);
