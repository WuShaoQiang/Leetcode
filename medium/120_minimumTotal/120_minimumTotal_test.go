package minimumTotal

func TestminimumTotal(t *testing.T){
tests:=map[string]struct{
input
want
}{
"Test 1" :{
input:,
want:,
},
}
for name,tc:= range tests{
t.Run(name,func(t *testing.T){
got:=minimumTotal(tc.input)
diff:=cmp.Diff(got,tc.want)
if diff!=""{
t.Fatalf("%v\n",diff)
}
})
}
}