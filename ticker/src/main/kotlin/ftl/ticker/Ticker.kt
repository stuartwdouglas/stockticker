package ftl.ticker

import xyz.block.ftl.Cron
import xyz.block.ftl.Export
import xyz.block.ftl.SinglePartitionMapper
import xyz.block.ftl.Topic
import xyz.block.ftl.WriteableTopic
import kotlin.random.Random

data class Price(val code: String, val price: Double, val timestamp: Long = System.currentTimeMillis())

@Export
@Topic(name = "prices", partitions = 1)
interface PriceTopic : WriteableTopic<Price, SinglePartitionMapper>

@Export
@Cron("1s")
fun tick(priceTopic: PriceTopic) {
    val p = Price("XYZ", Random.nextDouble() * 10)
    priceTopic.publish(p)
}
