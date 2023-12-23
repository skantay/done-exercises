#! /bin/bash
export INT_NUM=$(grep "SEE INTERVIEW" streets/Buckingham_Place | cut -d "#" -f2)
echo $INT_NUM
head "interviews/interview-$INT_NUM"
echo $MAIN_SUSPECT
