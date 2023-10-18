use std::fs;
use std::io::{prelude::*, BufReader};

macro_rules! die {
    ($code:expr, $($args:tt)*) => {
        ::std::eprintln!($($args)*);
        ::std::process::exit($code)
    }
}

fn main() {
    let Some(infile) = std::env::args().skip(1).next() else {
        die!(1, "missing input file");
    };
    //
}

struct ProtoFile {
    //
}

struct ProtoMessage {
}

struct ProtoComment {
}

#[derive(Clone, Copy, PartialEq, Eq)]
enum ProtoScalar {
    Double,
    Float,
    Int32,
    Int64,
    Uint32,
    Uint64,
    Sint32,
    Sint64,
    Fixed32,
    Fixed64,
    Sfixed32,
    Sfixed64,
    Bool,
    String,
    Bytes,
}

impl ProtoScalar {
    fn to_julia_type(self) -> &'static str {
        use Self::*;
        match self {
            Double => "Float64",
            Float => "Float32",
            Int32 => "Int32",
            Int64 => "Int64",
            Uint32 => "UInt32",
            Uint64 => "UInt64",
            Sint32 => "Int32",
            Sint64 => "Int64",
            Fixed32 => "UInt32",
            Fixed64 => "UInt64",
            Sfixed32 => "Int32",
            Sfixed64 => "Int64",
            Bool => "Bool",
            String => "String",
            Bytes => "Vector{UInt8}",
        }
    }
}
