
enum TSEnum {
    Foo,
    Bar,
    Baz
}

type Foo = {
    bar? : string
}

function doMe(foo: Foo): Boolean {
    if (foo.bar){
        return true;
    } else{
        return false
    }
}