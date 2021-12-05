mod blockchain_client;

use crate::blockchain_client::{BSCBlockchainClient, BlockHeader, BlockNumber, BlockchainClient};
use anyhow::Result;
use clap::{App, Arg, ArgMatches};
use std::time::Duration;

struct Checker {
    latest_block_header: BlockHeader,
}

#[derive(Debug)]
enum CheckerResult {
    Success,
    BlockNumberBack {
        previous_block_header: BlockHeader,
        new_block_number: BlockNumber,
    },
    ParentDidnotMatch {
        previous_block_header: BlockHeader,
        new_block_header: BlockHeader,
    },
}

impl Checker {
    pub fn new(block_header: BlockHeader) -> Self {
        Self {
            latest_block_header: block_header,
        }
    }

    pub fn check_block_number_backward(&self, new_block_number: BlockNumber) -> CheckerResult {
        if new_block_number < self.latest_block_header.number {
            CheckerResult::BlockNumberBack {
                previous_block_header: self.latest_block_header.clone(),
                new_block_number: new_block_number,
            }
        } else {
            CheckerResult::Success
        }
    }

    pub fn check_parent_hash(&self, new_block_header: BlockHeader) -> CheckerResult {
        if new_block_header.parent_hash != self.latest_block_header.hash {
            CheckerResult::ParentDidnotMatch {
                previous_block_header: self.latest_block_header.clone(),
                new_block_header,
            }
        } else {
            CheckerResult::Success
        }
    }

    pub fn update(&mut self, block_header: BlockHeader) {
        self.latest_block_header = block_header;
    }
}

struct Notifier {}

#[derive(Debug)]
struct Notification {
    message: String,
}

impl From<CheckerResult> for Notification {
    fn from(checker_result: CheckerResult) -> Self {
        Notification {
            message: format!("{:?}", checker_result),
        }
    }
}

impl Notifier {
    fn new() -> Notifier {
        Notifier {}
    }

    fn notify(&mut self, notification: Notification) {
        println!("notification: {:?}", notification)
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
    let node_url = "https://bsc-dataseed.binance.org";
    let blockchain_client = BSCBlockchainClient::new(node_url)?;
    let mut latest_header = blockchain_client.get_latest_block_header().await?;
    let mut checker = Checker::new(latest_header.clone());
    let mut notifier = Notifier::new();

    loop {
        tokio::time::sleep(Duration::from_secs(1)).await;

        let latest_block_number = blockchain_client.get_latest_number().await;
        let latest_block_number = match latest_block_number {
            Ok(block_number) => block_number,
            Err(err) => {
                eprintln!("Error getting latest block number: {}", err);
                continue;
            }
        };
        let check_result = checker.check_block_number_backward(latest_block_number);
        if !matches!(check_result, CheckerResult::Success) {
            notifier.notify(check_result.into());
            continue;
        };

        let next_block_header = blockchain_client
            .get_block_header(latest_header.number + 1)
            .await;
        let next_block_header = match next_block_header {
            Ok(Some(block_header)) => block_header,
            Ok(None) => {
                continue;
            }
            Err(err) => {
                eprintln!("failed to get next block header: {}", err);
                continue;
            }
        };
        let check_result = checker.check_parent_hash(next_block_header.clone());
        if !matches!(check_result, CheckerResult::Success) {
            notifier.notify(check_result.into())
        }
        checker.update(next_block_header.clone());
        latest_header = next_block_header;
        println!("new block {}: {}", latest_header.number, latest_header.hash);
    }

    Ok(())
}

fn clap_get_matches() -> ArgMatches<'static> {
    let matches = App::new("reorg checker")
        .version("0.1.0")
        .author("Park Juhyung <majecty@gmail.com>")
        .about("check reorganization and send notification")
        .arg(
            Arg::with_name("chain")
                .short("c")
                .long("chain")
                .value_name("BLOCKCHAIN_NAME")
                .help("like ethereum, klaytn, binance-smart-chain")
                .takes_value(true)
                .required(true),
        )
        .arg(
            Arg::with_name("network")
                .short("n")
                .long("network")
                .value_name("NETWORK_NAME")
                .help("like testnet or mainnet")
                .takes_value(true)
                .required(true),
        )
        .get_matches();
    matches
}
