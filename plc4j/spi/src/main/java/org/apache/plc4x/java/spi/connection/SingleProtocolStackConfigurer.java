/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.apache.plc4x.java.spi.connection;

import io.netty.buffer.ByteBuf;
import io.netty.channel.ChannelHandler;
import io.netty.channel.ChannelPipeline;
import org.apache.plc4x.java.spi.InstanceFactory;
import org.apache.plc4x.java.spi.Plc4xNettyWrapper;
import org.apache.plc4x.java.spi.Plc4xProtocolBase;
import org.apache.plc4x.java.spi.generation.Message;

import java.util.function.Consumer;
import java.util.function.Function;

/**
 * Builds a Protocol Stack.
 */
public class SingleProtocolStackConfigurer<BASE_PACKET_CLASS extends Message> implements ProtocolStackConfigurer<BASE_PACKET_CLASS> {

    private final Class<BASE_PACKET_CLASS> basePacketClass;
    private final Class<? extends Plc4xProtocolBase<BASE_PACKET_CLASS>> protocolClass;
    private final Class<? extends Function<ByteBuf, Integer>> packetSizeEstimator;
    private final Class<? extends Consumer<ByteBuf>> corruptPacketRemover;

    /** Only accessible via Builder */
    SingleProtocolStackConfigurer(Class<BASE_PACKET_CLASS> basePacketClass, Class<? extends Plc4xProtocolBase<BASE_PACKET_CLASS>> protocol,
                                  Class<? extends Function<ByteBuf, Integer>> packetSizeEstimator,
                                  Class<? extends Consumer<ByteBuf>> corruptPacketRemover) {
        this.basePacketClass = basePacketClass;
        this.protocolClass = protocol;
        this.packetSizeEstimator = packetSizeEstimator;
        this.corruptPacketRemover = corruptPacketRemover;
    }

    public static <BPC extends Message> SingleProtocolStackBuilder<BPC> builder(Class<BPC> basePacketClass) {
        return new SingleProtocolStackBuilder<>(basePacketClass);
    }

    private ChannelHandler getMessageCodec(InstanceFactory instanceFactory) {
        ReflectionBasedIo<BASE_PACKET_CLASS> io = new ReflectionBasedIo<>(basePacketClass);
        return new GeneratedProtocolMessageCodec<>(basePacketClass, io, io,
            packetSizeEstimator != null ? instanceFactory.createInstance(packetSizeEstimator) : null,
            corruptPacketRemover != null ? instanceFactory.createInstance(corruptPacketRemover) : null);
    }

    /** Applies the given Stack to the Pipeline */
    @Override
    public Plc4xProtocolBase<BASE_PACKET_CLASS> apply(InstanceFactory factory, ChannelPipeline pipeline) {
        pipeline.addLast(getMessageCodec(factory));
        Plc4xProtocolBase<BASE_PACKET_CLASS> protocol = factory.createInstance(protocolClass);
        Plc4xNettyWrapper<BASE_PACKET_CLASS> context = new Plc4xNettyWrapper<>(pipeline, protocol, basePacketClass);
        pipeline.addLast(context);
        return protocol;
    }

}
