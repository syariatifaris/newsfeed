INSERT INTO public.topic (name,is_valid) VALUES
  ('politik',1)
  ,('mancanegara',1)
  ,('teknologi',1);

INSERT INTO public.news (html_content,status,title,is_valid) VALUES
  ('Pertemuan presiden Kim Jong Un dan ..','published','Pertemuan Kim Jong Un dan Donald Trump di Singapore',1)
  ,('Setelah terjadi.....','published','Demo Presiden Korea Selatan',1)
  ,('Setelah terjadi.....','published','Demo Presiden Korea Selatan',1)
  ,('Setelah terjadi.....','published','Demo Presiden Korea Selatan',1)
  ,('Setelah terjadi.....','published','Demo Presiden Korea Selatan',1)
  ,('Setelah terjadi.....','published','Demo Presiden Korea Selatan',1)
  ,('Setelah terjadi.....','published','Demo Presiden Korea Selatan',1)
  ,('Setelah terjadi.....','published','Demo Presiden Korea Selatan',1)
  ,('Setelah terjadi.....','published','Demo Presiden Korea Selatan',1)
  ,('Samsung mengeluarkan','published','Flagship terbaru dari Samsung',1);

INSERT INTO public.topic_news (topic_id,news_id,is_valid) VALUES
  (1,11,1)
  ,(2,11,1)
  ,(1,12,1)
  ,(2,12,1)
  ,(3,13,1);