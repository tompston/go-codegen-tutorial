{{ $day := .Day -}}
{{ $day_concat  := printf "%s%s" "Concatenated" $day -}}

// This text file is generated

Today is {{$day}}

This is a new variable -> {{$day_concat}}