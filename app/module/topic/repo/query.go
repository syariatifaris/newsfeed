package repo

const getIDsByNames = `SELECT topic_id FROM topic WHERE name IN (?)`

const insertTopicNews = `INSERT INTO topic_news(topic_id, news_id) VALUES(?, ?)`
