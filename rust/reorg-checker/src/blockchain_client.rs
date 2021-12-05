use anyhow::Result;
use async_trait::async_trait;
use web3::transports::Http;
use web3::types::BlockNumber as Web3BlockNumber;
use web3::types::{BlockId, U64};
use web3::Result as Web3Result;
use web3::Web3;

pub type BlockNumber = u64;
pub type Hash = String;

#[derive(Debug, Clone)]
pub struct BlockHeader {
    pub number: BlockNumber,
    pub hash: Hash,
    pub parent_hash: Hash,
}

#[async_trait]
pub trait BlockchainClient {
    async fn get_latest_number(&self) -> Result<BlockNumber>;
    async fn get_block_header(&self, block_number: BlockNumber) -> Result<Option<BlockHeader>>;
    async fn get_latest_block_header(&self) -> Result<BlockHeader>;
}

pub struct BSCBlockchainClient {
    web3: Web3<Http>,
}

impl BSCBlockchainClient {
    pub fn new(node_url: &str) -> Result<BSCBlockchainClient> {
        let transport = web3::transports::Http::new(node_url)?;
        let web3 = web3::Web3::new(transport);
        Ok(BSCBlockchainClient { web3 })
    }
}
#[async_trait]
impl BlockchainClient for BSCBlockchainClient {
    async fn get_latest_number(&self) -> Result<BlockNumber> {
        let result: Web3Result<U64> = self.web3.eth().block_number().await;
        Ok(result?.as_u64())
    }

    async fn get_block_header(&self, block_number: BlockNumber) -> Result<Option<BlockHeader>> {
        let block_id: BlockId = U64::from(block_number).into();
        let result = self.web3.eth().block(block_id).await?;
        Ok(result.map(|block| BlockHeader {
            number: block.number.unwrap_or_default().as_u64(),
            hash: block.hash.unwrap_or_default().to_string(),
            parent_hash: block.parent_hash.to_string(),
        }))
    }

    async fn get_latest_block_header(&self) -> Result<BlockHeader> {
        let block = self
            .web3
            .eth()
            .block(BlockId::Number(Web3BlockNumber::Latest))
            .await?
            .unwrap();
        Ok(BlockHeader {
            number: block.number.unwrap_or_default().as_u64(),
            hash: block.hash.unwrap_or_default().to_string(),
            parent_hash: block.parent_hash.to_string(),
        })
    }
}
