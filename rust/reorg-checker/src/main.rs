use clap::{Arg, App};

#[tokio::main]
async fn main() -> Result<(), ()> {
   let matches = App::new("reorg checker")
       .version("0.1.0")
       .author("Park Juhyung <majecty@gmail.com>")
       .about("check reorganization and send notification")
       .arg(Arg::with_name("chain")
           .short("c")
           .long("chain")
           .value_name("BLOCKCHAIN_NAME")
           .help("like ethereum, klaytn, binance-smart-chain")
           .takes_value(true))
       .arg(Arg::with_name("network")
           .short("n")
           .long("network")
           .value_name("NETWORK_NAME")
           .help("like testnet or mainnet")
           .takes_value(true))
       .get_matches();

    let config = matches.value_of("chain").unwrap();
    println!("config {}", config);
    let network = matches.value_of("network").unwrap();
    println!("network {}", network);

    println!("Hello, world!");
    Ok(())
}