package org.eientei.videostreamer.mp4.boxes;

import io.netty.buffer.ByteBuf;
import org.eientei.videostreamer.mp4.Mp4BoxFull;
import org.eientei.videostreamer.mp4.Mp4Context;

/**
 * Created by Alexander Tumin on 2016-10-23
 */
public class Mp4StszBox extends Mp4BoxFull {
    public Mp4StszBox(Mp4Context context) {
        super("stsz", context, 0, 0);
    }

    @Override
    protected void fullWrite(ByteBuf out) {
        out.writeInt(0); // sample size
        out.writeInt(0); // sample countx
    }
}
