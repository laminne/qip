-- CreateTable
CREATE TABLE "User" (
    "id" TEXT NOT NULL,
    "handle" TEXT NOT NULL,
    "nickName" TEXT NOT NULL,
    "role" INTEGER NOT NULL,
    "bio" TEXT NOT NULL,
    "headerImageURL" TEXT NOT NULL,
    "iconImageURL" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "isLocalUser" BOOLEAN NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "apDataID" TEXT NOT NULL,
    "serverId" TEXT,

    CONSTRAINT "User_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "UserAPData" (
    "id" TEXT NOT NULL,
    "inboxURL" TEXT NOT NULL,
    "outboxURL" TEXT NOT NULL,
    "followersURL" TEXT NOT NULL,
    "followingURL" TEXT NOT NULL,
    "publicKey" TEXT NOT NULL,
    "privateKey" TEXT,

    CONSTRAINT "UserAPData_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "UserFollowEvent" (
    "followerID" TEXT NOT NULL,
    "followingID" TEXT NOT NULL,

    CONSTRAINT "UserFollowEvent_pkey" PRIMARY KEY ("followingID","followerID")
);

-- CreateTable
CREATE TABLE "Server" (
    "id" TEXT NOT NULL,
    "host" TEXT NOT NULL,
    "softwareName" TEXT NOT NULL,
    "softwareVersion" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "maintainer" TEXT NOT NULL,
    "maintainerEmail" TEXT NOT NULL,
    "iconURL" TEXT NOT NULL,
    "faviconURL" TEXT NOT NULL,

    CONSTRAINT "Server_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Post" (
    "id" TEXT NOT NULL,
    "text" TEXT NOT NULL,
    "visibility" INTEGER NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "authorID" TEXT,

    CONSTRAINT "Post_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Media" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "type" TEXT NOT NULL,
    "md5Sum" TEXT NOT NULL,
    "size" INTEGER NOT NULL,
    "isSensitive" BOOLEAN NOT NULL,
    "blurhash" TEXT NOT NULL,
    "url" TEXT NOT NULL,
    "thumbnailURL" TEXT NOT NULL,
    "cached" BOOLEAN NOT NULL,
    "authorID" TEXT,
    "postID" TEXT,

    CONSTRAINT "Media_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "User_apDataID_key" ON "User"("apDataID");

-- CreateIndex
CREATE UNIQUE INDEX "Server_host_key" ON "Server"("host");

-- AddForeignKey
ALTER TABLE "User" ADD CONSTRAINT "User_apDataID_fkey" FOREIGN KEY ("apDataID") REFERENCES "UserAPData"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "User" ADD CONSTRAINT "User_serverId_fkey" FOREIGN KEY ("serverId") REFERENCES "Server"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "UserFollowEvent" ADD CONSTRAINT "UserFollowEvent_followerID_fkey" FOREIGN KEY ("followerID") REFERENCES "User"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "UserFollowEvent" ADD CONSTRAINT "UserFollowEvent_followingID_fkey" FOREIGN KEY ("followingID") REFERENCES "User"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Post" ADD CONSTRAINT "Post_authorID_fkey" FOREIGN KEY ("authorID") REFERENCES "User"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Media" ADD CONSTRAINT "Media_authorID_fkey" FOREIGN KEY ("authorID") REFERENCES "User"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Media" ADD CONSTRAINT "Media_postID_fkey" FOREIGN KEY ("postID") REFERENCES "Post"("id") ON DELETE SET NULL ON UPDATE CASCADE;
