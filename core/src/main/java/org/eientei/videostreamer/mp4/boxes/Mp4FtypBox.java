package org.eientei.videostreamer.mp4.boxes;

import io.netty.buffer.ByteBuf;
import org.eientei.videostreamer.mp4.Mp4Box;
import org.eientei.videostreamer.mp4.Mp4RemuxerHandler;

/**
 * Created by Alexander Tumin on 2016-10-22
 */
public class Mp4FtypBox extends Mp4Box {
    private final String brand;
    private final int version;
    private final String[] supbrands;

    public Mp4FtypBox(Mp4RemuxerHandler context, String brand, int version, String... supbrands) {
        super("ftyp", context);
        this.brand = brand;
        this.version = version;
        this.supbrands = supbrands;
    }

    @Override
    protected void doWrite(ByteBuf out) {
        out.writeBytes(brand.getBytes());
        out.writeInt(version);
        for (String b : supbrands) {
            out.writeBytes(b.getBytes());
        }
    }
}
