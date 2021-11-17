mod blockchain_client;

use std::time::Duration;
use clap::{Arg, App, ArgMatches};
use anyhow::Result;
use crate::blockchain_client::{BlockchainClient, BlockHeader, BlockNumber, BSCBlockchainClient};

struct Checker {
    latest_block_header: BlockHeader,

    new_block_number: Option<BlockNumber>,
    new_block_header: Option<BlockHeader>,
}

enum CheckerResult {
    BlockNumberBack {
        previous_block_header: BlockHeader,
        new_block_header: BlockHeader,
    },
    ParentDidnotMatch {
        previous_block_header: BlockHeader,
        new_block_header: BlockHeader,
    },
}

impl Checker {
    fn new(block_header: BlockHeader) -> Self {
        Self {
            latest_block_header: block_header,
            new_block_header: Option::None,
            new_block_number: Option::None,
        }
    }

    fn check_block_number_backward(&self, new_block_number: BlockNumber) -> CheckerResult {
        if new_block_number < self.latest_block_header.number {}
        not_implemented!()
    }

    fn check_parent_hash(&self, new_block_header: BlockHeader) {
        if (new_block_header)
    }

    fn update(&mut self, block_header: BlockHeader) {
        self.latest_block_number = block_header.number;
        self.latest_block_header = block_header;
    }
}

#[tokio::main]
async fn main() -> Result<()> {
    let matches = clap_get_matches();

    let config = matches.value_of("chain").unwrap();
    println!("config {}", config);
    let network = matches.value_of("network").unwrap();
    println!("network {}", network);


    let node_url = "https://data-seed-prebsc-1-s1.binance.org:8545/";
    let blockchain_client = BSCBlockchainClient::new(node_url)?;
    let latest_header = blockchain_client.get_latest_block_header().await?;
    let checker = Checker::new(latest_header);

    loop {
        tokio::time::sleep(Duration::from_secs(1));

        let latest_block_number = blockchain_client.get_latest_number().await?;
        checker.check_block_number_backward(latest_block_number)
        let latest_block_header = blockchain_client.get_block_header(latest_header.number).await?;
        checker.check_block_number_backward()
    }


    let latest_block_number = blockchain_client.get_latest_number();

    println!("Hello, world!");
    Ok(())
}

fn clap_get_matches() -> ArgMatches {
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
    matches
}