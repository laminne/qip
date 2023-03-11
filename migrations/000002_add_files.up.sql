ALTER TABLE "files"
    ADD COLUMN uploaderID varchar NOT NULL;
ALTER TABLE "files"
    ADD COLUMN postID varchar;

ALTER TABLE "files"
    ADD FOREIGN KEY (uploaderID) REFERENCES users (id);
ALTER TABLE "files"
    ADD FOREIGN KEY (postID) REFERENCES posts (id);

ALTER TABLE "users"
    ADD FOREIGN KEY (headerImageID) REFERENCES files (id);
ALTER TABLE "users"
    ADD FOREIGN KEY (iconImageID) REFERENCES files (id);


