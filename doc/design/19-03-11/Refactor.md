# 从一个方法开始

    /**
     * @return The density object, be like `int[][]`
     * @param session The session to get density
     * @param viewport Get density in this viewport
     * @param condition Where condition of the query (in fact, it's just sql where statement)
     */
     Density getHotspotDensity(MdpSession session, Rectangle viewport, QueryCondition condition) 
    
这个方法来自于CHD `SessionResultService` 接口，大家看看有什么问题呢？

# 以发展的眼光看问题

这个标题是唯物辩证法的核心方法论之一，对我们的生活有着指导意义。设想你现在要给家里置办一口锅，有哪些方面需要考虑呢？

- 炒菜好不好吃
- 炒完好不好洗
- 能不能拿来烙饼

这三点对应到我们的代码上就是常说的：可用性，可维护性，可拓展性。那么我们再来检阅我们上面的代码。

- 可用么？当然可用，不然就报bug了，当然性能暂时不做考量。
- 可维护么？乍一看不知道什么叫可维护，注意用发展的眼光看问题。答案是不好维护，这一点后面会说。
- 可拓展么？不！

# 变与不变

变化是这个世界的主题，尤其是提需求的同志们深谙其道。且不说两个release，QA1和QA2之间都是瞬息万变。你看，这不是新的需求来了

- 添加一个pattern选项，用户可以看指定pattern的hotspot分布
- 添加一个resolution选项，用户可以指定看10x10还是50x50的图

对于的一个需求，你也许会说：“简单，加上一行呗”

    condition.addWhere("pfc_samples.adel_pattern_id = " + getPatternId());
    
姑且算你过关，先按下不表。那么第二个问题呢？“更简单了，加个参数呗”

    Density getHotspotDensity(MdpSession session, Rectangle viewport, QueryCondition condition, int resolution)
     
看起来倒是简单，但是迎接你的将是无穷无尽的error。因为所有的实现以及上下游调用都要改，你要花15分钟甚至更多时间把error清理掉再开始真正的修改。

然而这15分钟还不是最重要的，因为以此法行事，这痛的15分钟将无穷往复。更重要的是，现在我们的代码上下游全由我们自己维护掌控，如果我们的代码存在客户代码，那么任何API的修改都将是致命的。回看本节标题，世界总在变，那不变又是谁呢？唯有API永恒。这就是*SOLID*法则中的**O**，开闭原则(Open-Closed Principle)

回到我们的例子上，我们该怎么重构呢？一个经典的做法是，永远只使用可控的封装对象。

    public class DensityParam{
      Rectangle viewport;
      int resolution;
      
      public DensityParam(Rectangle, int){
        ...
      }
    }
    
    Density getHotspotDensity(MdpSession session, QueryCondition condition, DensityParam param)
    
这样一来，对于新的需求我们只需要再param类上加一个field就再也不用修改上下游所有接口及实现了。但是这里还有两个问题

- 其实我们还是没有满足开闭原则，拓展依然存在额外成本，你能发现么？稍后揭晓。
- 为什么没有把condition也放进param里？       因为我们要再下一节干掉他！

# 不该知道的别知道

设想你(Consumer Module)拦下一辆出租车，跟司机(DriveService)说去海王大厦(Expect param `Position target`)。结果司机说可以，但是你要给我指路(Actual param `RouteInfo[] routes`)。那么问题来了

- 你可能根本就不知道怎么走，你停在路边查导航，说不定你还得先下个地图APP(extra knowledge, extra effort)
- 你知道怎么走，然后快到了发现创业路封路了(bad expandability)

正确的做法应该是司机(DriveService)只接收真正的用户输入(目的地)，然后把这一请求委托给自己的脑袋(InMemoryRouteDAO)亦或是导航(ThirdPartyRouteDAO)。

回到我们的例子，`QueryCondition`这一次参数无情地入侵了DAO的领地。这对上层下层都造成了巨大的伤害

- 对于GUI的Maintainer来说，他非常痛苦地要了解DAO的具体实现，用的是MySQL，有哪些表哪些列哪些限制。
- 对于DAO的Maintainer来说，他要非常小心谨慎防止破坏上层访问，一旦加入任何改动，都要通知GUI重新测试评估。

由于参数的入侵，所谓的结构性解耦已经形同虚设。社会分工不复存在，大家都是“全才”，互相束缚手脚，牵一发而动全身。这也就是为什么我在前面说这个方法可维护性很差。

那么如何重构呢？很简单，不该知道的别知道，能不知道就不知道。客户如何输入我们就如实反映给下游，不要僭越做额外的工作。

    public class DensityParam{
      Rectangle viewport;
      int patternId;
      int resolution;
      
      public DensityParam(Rectangle, int, int){
        ...
      }
    }
    
    Density getHotspotDensity(MdpSession session, DensityParam param)

下层需要知道的仅仅是通过哪个pattern来filter，而不是你自作主张组合的sql语句。这样才能做到真真的解耦，把风险保持在可控范围内。

# 开闭到底

在前面我们根据开闭原则重构了我们的接口，引入了`DensityParam`对象，但是我们做的还不够彻底。如果你曾经读过*Effective Java*，甚至只翻开前几页，你就会看到

> Item 1: Consider static factory methods instead of constructors
> Item 2: Consider a builder when faced with many constructor parameters

这下你应该意识到了，虽然接口不在会被破坏，但是param的构造函数依然被暴露且被修改，那么所有调用者也都必须与时俱进。这其实并不是我们所期望的，任何新加的参数都应该有默认值，以让原有调用都仍然兼容。你也许会简单的给构造器加上一个重载

     public DensityParam(Rectangle, int)
      
     public DensityParam(Rectangle, int, int)

这种方法当然可行，但是随着时间推移，你的param类有了10个重载构造器，每个都有一大串参数以及默认值，调用者得仔细辨识每一个构造器防止调错了。

正确的做法应该是采用Builder模式，来创建真正的开闭兼容的API

    DensityParam.builder()
                .viewport(somewhere.getViewport())
                .patternId(someId)
                .build(); // resolution default value is 100
                
采用了Builder模式后，不仅获得了强兼容性，也增强了可读性，这可比一长串参数的构造器好理解多了。

[Lombok @Builder Sample](../../../src/main/java/xdean/share/pieces/BuilderDemo.java)

# 結束

至此，我们的重构结束了，经此一役，我们

- 解決了可维护性缺陷，解耦了UI和DAO
  - UI再也不必了解DB结构和SQL语句
  - DAO再也不用担心改变内部实现导致bug
- 解决了拓展性缺陷，再也不必为了增加一个参数而进行从上到下的大换血
- 解决了可读性缺陷，调用者不必阅读源码或doc，只从API上就能了解到必要信息