#!/usr/bin/python

import os, shutil
import subprocess
import os.path
from datetime import datetime
import argparse

######################## Functions #########################
class PhotoCat(object):

    def __init__(self, args = {}):
        self.sourcedir = args.source_path
        self.destDir = args.dest_path
        self.suffix = args.suffix
        self.organize_files()

    def photoDate(self, f):
      "Return the date/time on which the given photo was taken."

      cDate = subprocess.check_output(['sips', '-g', 'creation', f]).decode("utf-8")
      cDate = cDate.split('\n')[1].lstrip().split(': ')[1] 
      return datetime.strptime(cDate, "%Y:%m:%d %H:%M:%S")


    ###################### Main program ########################
    def organize_files(self):
        # The format for the new file names.
        fmt = "%Y-%m-%d_%H.%M.%S"

        # The problem files.
        problems = []

        # Get all the JPEGs in the source folder.
        photos = os.listdir(self.sourcedir)
        photos = [ x for x in photos if x[-4:] == '.jpg' or x[-4:] == '.JPG' ]

        # Copy photos into year and month subfolders. Name the copies according to
        # their timestamps. If more than one photo has the same timestamp, add
        # suffixes 'a', 'b', etc. to the names. 
        for photo in photos:
          original = self.sourcedir + '/' + photo
          suffix = self.suffix
          try:
            pDate = self.photoDate(original)
            yr = pDate.year
            mo = pDate.month
            newname = pDate.strftime(fmt)            
            oldname = photo.split(".")[0]
            duplicate = self.destDir + '/%04d/%02d/%s.jpg' % (yr, mo, newname)
            newname = pDate.strftime(fmt) + "_" + suffix + oldname
            new_dir = '%s/%04d/%02d' % (self.destDir, yr, mo)
            if not os.path.exists(new_dir):
                print("creating" + new_dir)
                os.makedirs(new_dir)
            duplicate = self.destDir + '/%04d/%02d/%s.jpg' % (yr, mo, newname)
            shutil.copy2(original, duplicate)
          except ValueError:
            dir = self.destDir + "/undated/"
            if not os.path.exists(dir):
                os.makedirs(dir)
            duplicate = dir + photo
            shutil.copy2(original, duplicate)
            problems.append(photo)

        # Report the problem files, if any.
        if len(problems) > 0:
          print("Problem files:")
          print("We've moved these bad boys over to undated")
          print("\n".join(problems))

if __name__ == "__main__":

    parser = argparse.ArgumentParser(
        formatter_class=argparse.ArgumentDefaultsHelpFormatter)

    parser.add_argument('--source_path', default=".",
        help="""This is the path to the pile of images""")
    parser.add_argument('--dest_path', default=".",
        help="""This is the path to put the output images""")
    parser.add_argument('--suffix', default="",
        help="""This is the suffix for all destination images""")

    args = parser.parse_args()
    p = PhotoCat(args)


