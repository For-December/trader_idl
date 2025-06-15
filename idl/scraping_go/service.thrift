namespace go scraping
namespace py scraping

// 1. 数据模型定义
struct PriceData {
  1: i64 timestamp,        // 时间戳（毫秒）
  2: double open,          // 开盘价
  3: double high,          // 最高价
  4: double low,           // 最低价
  5: double close,         // 收盘价
  6: double volume,        // 成交量
}

struct TextData {
  1: string content,       // 文本内容
  2: i64 timestamp,        // 发布时间
  3: string platform,      // 平台标识（twitter/reddit）
  4: optional string author,// 作者（可选）
}

// 2. 数据获取服务接口
service ScrapingService {
  string Echo(1: string message);

  // 获取BTC历史价格数据
  list<PriceData> GetBtcPrice(1: i64 startTime, 2: i64 endTime, 3: string interval),
  
  // 获取社交媒体文本数据
  list<TextData> GetSocialMediaData(1: string keyword, 2: i64 startTime, 3: i64 endTime, 4: i32 limit),
  
}
