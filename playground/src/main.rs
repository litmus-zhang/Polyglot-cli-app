use std::path::PathBuf;

use anyhow::{Result, anyhow, Context};

fn error_me(throw: bool) -> Result<()> {
    if throw {
        return Err(anyhow!("This is not true"));
    }
    std::fs::read(PathBuf::from("/foo")).context("Unable tp do this")?;
    return Ok(());
}
fn main() -> Result<(), usize> {
    let value = error_me(false)?;         
    return Ok(());
}
