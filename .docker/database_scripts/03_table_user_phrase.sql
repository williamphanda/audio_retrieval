-- Table: public.audio_retrieval_trx_user_phrases

-- DROP TABLE public.audio_retrieval_trx_user_phrases;

CREATE TABLE IF NOT EXISTS audio_retrieval_trx_user_phrases(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    phrase_id BIGINT NOT NULL,
    file_path varchar(255) NOT NULL,
    mime_type SMALLINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES audio_retrieval_trx_users (id) ON DELETE CASCADE,
    CONSTRAINT fk_phrase FOREIGN KEY (phrase_id) REFERENCES audio_retrieval_trx_phrases (id) ON DELETE CASCADE
);