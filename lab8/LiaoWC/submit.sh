if [ "$#" -ne 2 ]; then
  echo "Usage: ./golang_submit.sh lab_no attempt_no"
  exit 1
fi
git add . && git commit -m "lab$1"
git branch "LiaoWC-lab$1-attempt$2" && git checkout "LiaoWC-lab$1-attempt$2"
git push --set-upstream origin "LiaoWC-lab$1-attempt$2"
