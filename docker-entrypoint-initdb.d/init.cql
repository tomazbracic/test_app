CREATE KEYSPACE IF NOT EXISTS projekt1 WITH replication = 
{'class':'SimpleStrategy','replication_factor':'1'};

CREATE TABLE IF NOT EXISTS test ( id uuid PRIMARY KEY, cas timestamp );