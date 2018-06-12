package repo

const getAllNews = `SELECT 
						* FROM news`

const getAllNewsByStatus = `SELECT 
								* FROM news WHERE (status = $1 AND is_valid = 1)`

const getAllNewsByTopics = `SELECT 
								* FROM news 
										WHERE is_valid = 1 AND news_id IN(
								SELECT DISTINCT(tn.news_id) 
									FROM topic_news tn 
										INNER JOIN topic t ON tn.topic_id = t.topic_id 
											WHERE t.name IN (?))`

const insertNews = `INSERT INTO news(title, html_content, status) VALUES(?, ?, ?) RETURNING news_id`
