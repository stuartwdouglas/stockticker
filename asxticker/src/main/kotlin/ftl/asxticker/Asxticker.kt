package ftl.asxticker

import ftl.ticker.Price
import xyz.block.ftl.*
import kotlin.random.Random

@Export
@Topic(name = "asxPrices", partitions = 1)
interface AsxPrice : WriteableTopic<Price, SinglePartitionMapper>

@Cron("3s")
fun tick(asxPrice: AsxPrice) {
    var price = Random.nextDouble() * 10
    val p = Price("ASX", price, System.currentTimeMillis())
    asxPrice.publish(p)
}