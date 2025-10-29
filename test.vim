"vimi Standard Library
"this is a script containing basic functionality
"which helps making vimi a full fledged interpreter


com! -nargs=* Import call Import(<f-args>)

function! Import(lib)
    exec "source " . a:lib . ".vim"
endfunction

let strings = ['foo', 'bar', 'baz', 'quux']
for i in range(0, len(strings) - 1)
  let item = strings[i]
  echo i + 1 . ":" . item
endfor
