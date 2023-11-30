OUT=''
for CHARPOS in {1..8}
do 
    LETTER=$(cat input.txt | awk -v i=$CHARPOS '{print substr($1, i, 1)}' | sort | uniq -c | sort -nr | head -n1 | awk '{print $2}')
    OUT="$OUT$LETTER"
done 
echo $OUT

OUT=''
for CHARPOS in {1..8}
do 
    LETTER=$(cat input.txt | awk -v i=$CHARPOS '{print substr($1, i, 1)}' | sort | uniq -c | sort -n | head -n1 | awk '{print $2}')
    OUT="$OUT$LETTER"
done 
echo $OUT