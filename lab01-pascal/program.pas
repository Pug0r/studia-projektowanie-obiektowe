program HelloWorld;

const
  ArraySize = 50;

type
  IntArray = array[1..50] of Integer;

var
  GeneratedNumbers: IntArray;
  i: Integer;

procedure GenerateNumbers(var targetArray: IntArray);
var
  i: Integer;
begin
  Randomize;
  for i := 1 to ArraySize do
    targetArray[i] := Random(101); 
end;

procedure DisplayArray(t: IntArray);
var
  i: Integer;
begin
  for i := 1 to ArraySize do
    Write(t[i], ' ');
  Writeln;
end;

begin

  GenerateNumbers(GeneratedNumbers);
  Writeln('Wylosowane liczby:');
  DisplayArray(GeneratedNumbers);

end.