/*
  Warnings:

  - Made the column `authorID` on table `Media` required. This step will fail if there are existing NULL values in that column.
  - Made the column `postID` on table `Media` required. This step will fail if there are existing NULL values in that column.
  - Made the column `authorID` on table `Post` required. This step will fail if there are existing NULL values in that column.

*/
-- DropForeignKey
ALTER TABLE "Media" DROP CONSTRAINT "Media_authorID_fkey";

-- DropForeignKey
ALTER TABLE "Media" DROP CONSTRAINT "Media_postID_fkey";

-- DropForeignKey
ALTER TABLE "Post" DROP CONSTRAINT "Post_authorID_fkey";

-- AlterTable
ALTER TABLE "Media" ALTER COLUMN "authorID" SET NOT NULL,
ALTER COLUMN "postID" SET NOT NULL;

-- AlterTable
ALTER TABLE "Post" ALTER COLUMN "authorID" SET NOT NULL;

-- CreateTable
CREATE TABLE "Reaction" (
    "userId" TEXT NOT NULL,
    "postId" TEXT NOT NULL,

    CONSTRAINT "Reaction_pkey" PRIMARY KEY ("userId","postId")
);

-- AddForeignKey
ALTER TABLE "Post" ADD CONSTRAINT "Post_authorID_fkey" FOREIGN KEY ("authorID") REFERENCES "User"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Reaction" ADD CONSTRAINT "Reaction_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Reaction" ADD CONSTRAINT "Reaction_postId_fkey" FOREIGN KEY ("postId") REFERENCES "Post"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Media" ADD CONSTRAINT "Media_authorID_fkey" FOREIGN KEY ("authorID") REFERENCES "User"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Media" ADD CONSTRAINT "Media_postID_fkey" FOREIGN KEY ("postID") REFERENCES "Post"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
