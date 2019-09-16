use std::time::Duration;

use async_std::{
    prelude::*,
    task,
    io,
    net::TcpStream,
};

async fn get() -> io::Result<Vec<u8>> {
    let mut stream = TcpStream::connect("hoper.xyz:80").await?;
    stream.write_all(b"GET / HTTP/1.0\r\n\r\n").await?;

    let mut buf = vec![];

    io::timeout(Duration::from_secs(5), async {
        stream.read_to_end(&mut buf).await?;
        Ok(buf)
    }).await
}

fn main() {
    task::block_on(async {
        let raw_response = get().await.expect("request");
        let response = String::from_utf8(raw_response)
            .expect("utf8 conversion");
        println!("received: {}", response);
    });
}